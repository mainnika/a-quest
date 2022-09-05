package keys

var (
	PublicKey = []byte(`-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEqtyaUimtphsWdH6aRxsFMi/TXkff
lpM2JTy7A94ut9Gk2HhOF05hiJuwFQQUN78WNhauZnbU1lLWPKP5lJbZ7Q==
-----END PUBLIC KEY-----`)
	PrivateKey = []byte(`-----BEGIN EC PRIVATE KEY-----
MHcCAQEEILLczON2Ou3TVZGzNZfNiP5XZCRALoPGLCoHe3m3jFQhoAoGCCqGSM49
AwEHoUQDQgAEqtyaUimtphsWdH6aRxsFMi/TXkfflpM2JTy7A94ut9Gk2HhOF05h
iJuwFQQUN78WNhauZnbU1lLWPKP5lJbZ7Q==
-----END EC PRIVATE KEY-----`)
	Alg    = `ES256`
	Answer = []byte(`9834876dcfb05cb167a5c24953eba58c4ac89b1adf57f28f2f9d09af107ee8f0`)
)
