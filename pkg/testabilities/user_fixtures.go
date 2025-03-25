package testabilities

var Alice = User{
	Paymails: []Paymail{"alice@example.com"},

	// TODO: Make std private key as a base key (and recreate xPriv if needed)
	PrivKey: "xprv9s21ZrQH143K2stnKknNEck8NZ9buundyjYCGFGS31bwApaGp7oviHYVY9YAogmgvFC8EdsbsDReydnhDXrRrSXoNoMZczV9t4oPQREAmQ3",
}

var Bob = User{
	Paymails: []Paymail{"bob@example.com"},
	PrivKey:  "xprv9s21ZrQH143K3c3jkTBGijY5UsiHUdd3fSzRFD21c7cFduWX4m9nPrcuVrjQ76K234TFWgKF3f97HXggriPipBdhuof6bSvLGE74zCCgJds",
}

var Charlie = User{
	Paymails: []Paymail{"charlie@example.com"},
	PrivKey:  "xprv9s21ZrQH143K4b2JYp37EzEcK55k5wQDnXaH3ooi8oq9yHEj8TCWGuVnJoQvQVyHx3eyF6DyLDiteD6G5CLdKvTcG8QwiEZPyqUcvgmj9aK",
}
