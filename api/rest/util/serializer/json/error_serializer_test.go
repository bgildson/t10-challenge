package json

import (
	"fmt"
	"reflect"
	"testing"
)

func TestErrorSerializer(t *testing.T) {
	message := "error message"
	expected := []byte(fmt.Sprintf(`{"error":"%v"}`, message))
	result, err := ErrorSerializer(message)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("was expecting\n%#v\nbut returns\n%#v", expected, result)
	}

	if err != nil {
		t.Errorf("an unexpected error has occurred: %v", err)
	}
}
