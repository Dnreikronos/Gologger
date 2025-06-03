	file, err := os.Open(path)
	if err != nil {
		return err
	}
	_, err = file.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}

	wg.Add(1)
