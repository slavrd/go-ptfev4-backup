package helpers

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
)

const reqBodyString = `{ "password": "%s" }`

// PtfeBackup creates a backup of a PTFE instance using the provided access details
func PtfeBackup(host, token, pwd string, out io.Writer) error {

	var body = strings.NewReader(fmt.Sprintf(reqBodyString, pwd))
	var url = fmt.Sprintf("https://%s/_backup/api/v1/backup", strings.TrimSuffix(host, "/"))

	log.Printf("making request to %q", url)

	var req, err = http.NewRequest(http.MethodPost, url, body)

	if err != nil {
		return fmt.Errorf("error constructing http request: %v", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)

	if err != nil {
		return fmt.Errorf("error making http request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return fmt.Errorf("http request returned: %s", resp.Status)
	}

	written, err := io.Copy(out, resp.Body)

	if err != nil {
		return fmt.Errorf("error saving backup: %v", err)
	}

	log.Printf("saved backup file, size %d bytes\n", written)

	return nil
}

// PtfeRestore restores a backup to a PTFE instance using the provided access details
func PtfeRestore(host, token, pwd string, in io.Reader) error {

	var url = fmt.Sprintf("https://%s/_backup/api/v1/restore", strings.TrimSuffix(host, "/"))

	log.Printf("making request to %q", url)

	// using io.Pipe so that we don't load the file into memory to execute the upload
	r, w := io.Pipe()

	m := multipart.NewWriter(w)

	go func() {
		defer w.Close()
		defer m.Close()

		partConfig, err := m.CreateFormField("config")
		if err != nil {
			return
		}
		cr := strings.NewReader(fmt.Sprintf(reqBodyString, pwd))
		if _, err := io.Copy(partConfig, cr); err != nil {
			return
		}

		partSnapshot, err := m.CreateFormField("snapshot")
		if err != nil {
			return
		}

		if _, err := io.Copy(partSnapshot, in); err != nil {
			return
		}

	}()

	var req, err = http.NewRequest(http.MethodPost, url, r)

	if err != nil {
		return fmt.Errorf("error constructing http request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	req.Header.Set("Content-Type", m.FormDataContentType())

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)

	if err != nil {
		return fmt.Errorf("error making http request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return fmt.Errorf("http request returned: %s", resp.Status)
	}

	if err != nil {
		return fmt.Errorf("error restoring backup: %v", err)
	}

	return nil
}
