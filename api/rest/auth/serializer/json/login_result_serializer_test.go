package json

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLoginResultSerializer(t *testing.T) {
	token := "mytoken"
	expected := []byte(fmt.Sprintf(`{"token":"%v"}`, token))
	result, err := LoginResultSerializer(token)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("was expecting\n%#v\nbut returns\n%#v", expected, result)
	}

	if err != nil {
		t.Errorf("an unexpected error has occurred: %v", err)
	}
}
