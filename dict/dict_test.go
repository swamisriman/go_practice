package dict

import (
	"testing"
)

func TestDict(t *testing.T) {

	assertStringsEquality := func(t testing.TB, expected, actual string) {
		t.Helper()

		if expected != actual {
			t.Errorf("Expected: %q but Actual: %q", expected, actual)
		}
	}

	assertErrorsEquality := func(t testing.TB, expected, actual error) {
		t.Helper()

		if expected != actual {
			t.Errorf("Expected: %q but Actual: %q", expected, actual)
		}
	}

	assertKeyValuePairPresence := func(t testing.TB, dict Dictionary, key, value string) {
		t.Helper()

		observedVal, err := dict.Search(key)

		switch err {
		case DictKeyNotFoundErr:
			t.Fatal("error occurred while checking key-value pair presence", err)
		case nil:
			assertStringsEquality(t, value, observedVal)
		default:
			t.Errorf("expected key: %q not found!", key)
		}
	}

	dict := Dictionary{
		"testKey":  "testValue",
		"testKey1": "testValue1",
	}

	t.Run("Fetch a key's value - different cases", func(*testing.T) {
		val, err := dict.Search("testKey1")
		assertStringsEquality(t, val, "testValue1")
		assertErrorsEquality(t, err, nil)

		val, err = dict.Search("random")
		assertStringsEquality(t, val, "")
		assertErrorsEquality(t, err, DictKeyNotFoundErr)

	})

	t.Run("Add a key value pair - different cases", func(*testing.T) {
		err := dict.Add("testKey2", "testValue2")
		assertErrorsEquality(t, err, nil)
		assertKeyValuePairPresence(t, dict, "testKey2", "testValue2")

		err = dict.Add("testKey2", "randomVal")
		assertErrorsEquality(t, err, DictKeyCollisionError)
		assertKeyValuePairPresence(t, dict, "testKey2", "testValue2")
	})

	t.Run("Update a key value pair - different cases", func(*testing.T) {
		existingVal, err := dict.Update("testKey2", "newTestValue2")
		assertErrorsEquality(t, err, nil)
		assertKeyValuePairPresence(t, dict, "testKey2", "newTestValue2")
		assertStringsEquality(t, "testValue2", existingVal)

		existingVal, err = dict.Update("testKey3", "randomVal")
		assertErrorsEquality(t, err, DictKeyNotFoundErr)
		assertStringsEquality(t, "", existingVal)
	})

	t.Run("Delete a key value pair - different cases", func(*testing.T) {
		existingVal := dict.Delete("testKey")
		assertStringsEquality(t, "testValue", existingVal)
		_, err := dict.Search("testKey")
		assertErrorsEquality(t, err, DictKeyNotFoundErr)

		existingVal = dict.Delete("testKey")
		assertStringsEquality(t, "", existingVal)
	})
}
