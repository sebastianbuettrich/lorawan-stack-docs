func makeProtoc() (func(...string) error, *protocContext, error) {

	wd, err := os.Getwd()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get working directory: %w", err)
	}
	usr, err := user.Current()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get user: %w", err)
	}

	mountWD := filepath.ToSlash(filepath.Join(filepath.Dir(wd), "lorawan-stack"))
	return sh.RunCmd("docker", "run",
			"--rm",
			"--user", fmt.Sprintf("%s:%s", usr.Uid, usr.Gid),
			"--mount", fmt.Sprintf("type=bind,src=%s,dst=%s/api", filepath.Join(wd, "api"), mountWD),
			"--mount", fmt.Sprintf("type=bind,src=%s,dst=%s/go.thethings.network/lorawan-stack/v3/pkg/ttnpb", filepath.Join(wd, "pkg", "ttnpb"), protocOut),
			"--mount", fmt.Sprintf("type=bind,src=%s,dst=%s/doc", filepath.Join(wd, "doc"), mountWD),
			"--mount", fmt.Sprintf("type=bind,src=%s,dst=%s/v3/sdk/js", filepath.Join(wd, "sdk", "js"), mountWD),
			"-w", mountWD,
			fmt.Sprintf("%s:%s", protocName, protocVersion),
			fmt.Sprintf("-I%s", filepath.Dir(wd)),
		), &protocContext{
			WorkingDirectory: mountWD,
			UID:              usr.Uid,
			GID:              usr.Gid,
		}, nil
}

func withProtoc(f func(pCtx *protocContext, protoc func(...string) error) error) error {
	protoc, pCtx, err := makeProtoc()
	if err != nil {
		return errors.New("failed to construct protoc command")
	}
	return f(pCtx, protoc)
}

// HugoData generates Hugo data files.
func HugoData(context.Context) error {
	return withProtoc(func(pCtx *protocContext, protoc func(...string) error) error {
		if err := protoc(
			fmt.Sprintf("--hugodata_out=output_path=%[1]s:%[1]s", filepath.Join(pCtx.WorkingDirectory, "doc", "data")),
			fmt.Sprintf("%s/api/*.proto", pCtx.WorkingDirectory),
		); err != nil {
			return fmt.Errorf("failed to generate protos: %w", err)
		}
		return nil
	})
}

// HugoDataClean removes generated Hugo data files.
func HugoDataClean(context.Context) error {
	return filepath.Walk(filepath.Join("doc", "data", "api", "ttn.lorawan.v3"), func(path string, _ os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for _, ext := range []string{"enums.yml", "messages.yml", "services.yml"} {
			if strings.HasSuffix(path, ext) {
				if err := sh.Rm(path); err != nil {
					return err
				}
				return nil
			}
		}
		return nil
	})
}
