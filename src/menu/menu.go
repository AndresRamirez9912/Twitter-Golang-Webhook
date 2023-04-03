package menu

func Response(input string) (output string) {
	switch input {
	case "1":
		output = "El saldo de su cuenta finalizada en 123 es de 15.000.000"
	case "hola":
		output =
			`
		Bienvenido al Chat Bot del Banco Popular, por
		favor seleccione la opcion que necesita
		1. Consultar el saldo de mis cuentas
		2. Enviar dinero a una cuenta
		`
	default:
		output = "Por favor seleccione una opcion valida"
	}
	return
}
