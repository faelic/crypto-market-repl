package repl

import "testing"

func TestParseInputHelp(t *testing.T) {
	command, args := parseInput("/Help")

	if command != "/Help" {
		t.Errorf("expected /help as command but got %s", command)
	}

	if len(args) != 0 {
		t.Errorf("expected no args but got %v", len(args))
	}
}

func TestParseInputPriceWithCoin(t *testing.T) {
	command, args := parseInput("/price bitcoin")

	if command != "/price" {
		t.Errorf("expected /price as command but got %s", command)
	}

	if len(args) != 1 {
		t.Fatalf("expected 1 arg, got %d", len(args))
	}

	if args[0] != "bitcoin" {
		t.Errorf("expected bitcoin as arg but got %s", args[0])
	}
}

func TestParseInputTrimsExtraWhitespace(t *testing.T) {
	command, args := parseInput(" /price  bitcoin")

	if command != "/price" {
		t.Errorf("expected command to be /price but got %s", command)
	}

	if len(args) != 1 {
		t.Fatalf("expected length of args to be 1 but got %d", len(args))
	}

	if args[0] != "bitcoin" {
		t.Errorf("expected args to be bitcoin but got %s", args[0])
	}
}

func TestParseInputWithEmptyString(t *testing.T) {
	gotCommand, gotArgs := parseInput("")

	wantCommand := ""

	if gotCommand != wantCommand {
		t.Errorf("expected command to be %s but got %s", wantCommand, gotCommand)
	}

	wantLenArgs := 0

	if len(gotArgs) != wantLenArgs {
		t.Errorf("expected lenth of args to be %d but got %d", wantLenArgs, len(gotArgs))
	}
}

func TestParseInputWhitespaceOnly(t *testing.T) {
	gotCommand, gotArgs := parseInput(" ")

	wantCommand := ""
	if gotCommand != wantCommand {
		t.Errorf("expeced command to be %s but got %s", wantCommand, gotCommand)
	}

	wantLenArgs := 0

	if len(gotArgs) != wantLenArgs {
		t.Errorf("expected lenth of args to be %d but got %d", wantLenArgs, len(gotArgs))
	}
}
