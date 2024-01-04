package network

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
)

func HttpClient() http.Client {
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM([]byte(`-----BEGIN CERTIFICATE-----
MIIGPzCCBSegAwIBAgIQeoogINRa01nUqqWJFPM6TjANBgkqhkiG9w0BAQsFADCB
jzELMAkGA1UEBhMCR0IxGzAZBgNVBAgTEkdyZWF0ZXIgTWFuY2hlc3RlcjEQMA4G
A1UEBxMHU2FsZm9yZDEYMBYGA1UEChMPU2VjdGlnbyBMaW1pdGVkMTcwNQYDVQQD
Ey5TZWN0aWdvIFJTQSBEb21haW4gVmFsaWRhdGlvbiBTZWN1cmUgU2VydmVyIENB
MB4XDTIzMDYzMDAwMDAwMFoXDTI0MDczMDIzNTk1OVowHTEbMBkGA1UEAxMSd3d3
LmRlc2lnbnBpbG90LmFpMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA
lzPx9aUGNPIazHriHpcjRLD4ZBYwvacjtiPHryCzxC/9c7wigW8OsmbhVzBfiBwM
/HbjTT2d+/b8jZtR9wgtySkmRDoWCxBWWdyK9+6E27HgahzGWrN1er+WoUr0aZ+d
Yy2fArd8VBwK7jJ/gZ9dp3d/iN9lQk8ZpUBvwQ0NinbLCLXQ+719XFkgXi+Vh+1f
Xh976c467t12zsn9DBeNgVQhzGxuDqz0ZU220014Cys91eAY3XGzggZt14GhEqXi
FlKCKUVrecQymuZTBnd7vr1vyQjMn3LI1PZ7NTY6yKFffM/8xLgUNNCqWmOmCI0s
l7Yb0WZZjXV4A/2CkO5NlQIDAQABo4IDBjCCAwIwHwYDVR0jBBgwFoAUjYxexFSt
iuF36Zv5mwXhuAGNYeEwHQYDVR0OBBYEFL09d5yfyextMxyPREDRYwN6c6boMA4G
A1UdDwEB/wQEAwIFoDAMBgNVHRMBAf8EAjAAMB0GA1UdJQQWMBQGCCsGAQUFBwMB
BggrBgEFBQcDAjBJBgNVHSAEQjBAMDQGCysGAQQBsjEBAgIHMCUwIwYIKwYBBQUH
AgEWF2h0dHBzOi8vc2VjdGlnby5jb20vQ1BTMAgGBmeBDAECATCBhAYIKwYBBQUH
AQEEeDB2ME8GCCsGAQUFBzAChkNodHRwOi8vY3J0LnNlY3RpZ28uY29tL1NlY3Rp
Z29SU0FEb21haW5WYWxpZGF0aW9uU2VjdXJlU2VydmVyQ0EuY3J0MCMGCCsGAQUF
BzABhhdodHRwOi8vb2NzcC5zZWN0aWdvLmNvbTAtBgNVHREEJjAkghJ3d3cuZGVz
aWducGlsb3QuYWmCDmRlc2lnbnBpbG90LmFpMIIBgAYKKwYBBAHWeQIEAgSCAXAE
ggFsAWoAdwB2/4g/Crb7lVHCYcz1h7o0tKTNuyncaEIKn+ZnTFo6dAAAAYkM3870
AAAEAwBIMEYCIQC3s5Zo6EanrE/zvAuueUHzBzpJlQ2Vj5GBQfvIS2ny0wIhAK8P
N6U88f2kIeFkv+LS1tGjysBo4eERke1xC/9pSjb1AHYA2ra/az+1tiKfm8K7XGvo
cJFxbLtRhIU0vaQ9MEjX+6sAAAGJDN/PbQAABAMARzBFAiEA9jaXhAZywertCz15
R/sd3y8+9J6l32KnHqjdNRLpQysCIAuaKLAGStuLsBsvTogLXdEbRIu0CVZvrl89
JmZ93Z0dAHcA7s3QZNXbGs7FXLedtM0TojKHRny87N7DUUhZRnEftZsAAAGJDN/P
agAABAMASDBGAiEAvIYCcHeaOO+BczK7VRLfqjeTcKaruACfhqMbnNSTTzcCIQCl
es7JSoJDKfUDosUklVr9uAvum67Mypjm3rp1uLMtaTANBgkqhkiG9w0BAQsFAAOC
AQEAeUq6RG8N12jVqLdmSveh0GDWDVZhMVr25cBJzsuHs1d6O0fX0UHZJSV73P87
EAiWZp6XHwxNlMuqkj8omhtCan0DRETsAaPJ/tU0Z1yvyrAIebq93t9qJzcZBSN3
o3ZoBYzwrNPL5anLZZrWgm4NgTqHso65b9Xp2wiMR6xEOiODhHIqnJbYdTdztGfT
stasF9Eb6EhvHyLmeSB3Lv1iyblACAoM5ZzSuuXaS8eLjL23gc20/4MB+/QjOVT/
Yd+ZupxuE7lTbtw9ZVw1lpwjmqTC7pr8JYA+SHWO7EBinramXktGgDxHjzzq9yLH
WewWP2/LplVCsR+A27mKXQy63Q==
-----END CERTIFICATE-----
`))

	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}

	transport := &http.Transport{TLSClientConfig: tlsConfig}
	httpClient := http.Client{Transport: transport}

	return httpClient
}