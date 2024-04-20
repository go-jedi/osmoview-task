package changefield

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

type Data1 struct {
	Name string
	Type string
}

func TestChangeFieldData1(t *testing.T) {
	// Arrange
	tests := []struct {
		name  string
		input struct {
			Object    interface{}
			fieldName string
			newValue  interface{}
		}
		expected Data1
		err      error
	}{
		{
			name: "OK",
			input: struct {
				Object    interface{}
				fieldName string
				newValue  interface{}
			}{
				Object: Data1{
					Name: "Name",
					Type: "Type",
				},
				fieldName: "Name",
				newValue:  "NewName",
			},
			expected: Data1{
				Name: "NewName",
				Type: "Type",
			},
			err: nil,
		},
		{
			name: "ERROR (NOT_FOUND)",
			input: struct {
				Object    interface{}
				fieldName string
				newValue  interface{}
			}{
				Object: Data1{
					Name: "Name",
					Type: "Type",
				},
				fieldName: "Names",
				newValue:  "NewName",
			},
			expected: Data1{},
			err:      errors.New("поле Names не найдено"),
		},
		{
			name: "ERROR (NOT_RIGHT_TYPE)",
			input: struct {
				Object    interface{}
				fieldName string
				newValue  interface{}
			}{
				Object:    Data1{},
				fieldName: "Name",
				newValue:  "NewName",
			},
			expected: Data1{},
			err:      errors.New("obj должен быть структурой"),
		},
		{
			name: "ERROR (TYPE_MISMATCH)",
			input: struct {
				Object    interface{}
				fieldName string
				newValue  interface{}
			}{
				Object: Data1{
					Name: "Name",
					Type: "Type",
				},
				fieldName: "Name",
				newValue:  1,
			},
			expected: Data1{},
			err:      errors.New("типы не совпадают: поле Name имеет тип string, а значение имеет тип int"),
		},
	}
	//	 Act
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			const caseOk = "OK"
			const caseNotFound = "ERROR (NOT_FOUND)"
			const caseNotRightType = "ERROR (NOT_RIGHT_TYPE)"
			const caseTypeMismatch = "ERROR (TYPE_MISMATCH)"

			switch test.name {
			case caseOk:
				data1, ok := test.input.Object.(Data1)

				err := ChangeField(&data1, test.input.fieldName, test.input.newValue)

				// Assert
				assert.Equal(t, true, ok)
				assert.NoError(t, err)
				assert.Equal(t, err, test.err)
				assert.Equal(t, test.expected, data1)
			case caseNotFound:
				data1, ok := test.input.Object.(Data1)

				err := ChangeField(&data1, test.input.fieldName, test.input.newValue)

				// Assert
				assert.Equal(t, true, ok)
				assert.Error(t, err)
				assert.Equal(t, err, test.err)
			case caseNotRightType:
				err := ChangeField(&test.input.Object, test.input.fieldName, test.input.newValue)

				// Assert
				assert.Error(t, err)
				assert.Equal(t, err, test.err)
			case caseTypeMismatch:
				data1, ok := test.input.Object.(Data1)

				err := ChangeField(&data1, test.input.fieldName, test.input.newValue)

				// Assert
				assert.Equal(t, true, ok)
				assert.Error(t, err)
				assert.Equal(t, err, test.err)
			}
		})
	}
}

func BenchmarkChangeFieldData1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data1 := Data1{Name: "Name", Type: "Type"}
		err := ChangeField(&data1, "Name", "NewName")
		require.NoError(b, err)
	}
}

type Data2 struct {
	ID  string
	Arr []int
	Map map[string]interface{}
	Any interface{}
}

func TestChangeFieldData2(t *testing.T) {
	// Arrange
	tests := []struct {
		name  string
		input struct {
			Object    interface{}
			fieldName string
			newValue  interface{}
		}
		expected Data2
		err      error
	}{
		{
			name: "OK",
			input: struct {
				Object    interface{}
				fieldName string
				newValue  interface{}
			}{
				Object: Data2{
					ID:  "1",
					Arr: []int{1, 1, 1},
					Map: map[string]interface{}{
						"map": "map",
					},
					Any: "any",
				},
				fieldName: "ID",
				newValue:  "111",
			},
			expected: Data2{
				ID:  "111",
				Arr: []int{1, 1, 1},
				Map: map[string]interface{}{
					"map": "map",
				},
				Any: "any",
			},
			err: nil,
		},
		{
			name: "ERROR (NOT_FOUND)",
			input: struct {
				Object    interface{}
				fieldName string
				newValue  interface{}
			}{
				Object: Data2{
					ID:  "1",
					Arr: []int{1, 1, 1},
					Map: map[string]interface{}{
						"map": "map",
					},
					Any: "any",
				},
				fieldName: "Names",
				newValue:  "111",
			},
			expected: Data2{},
			err:      errors.New("поле Names не найдено"),
		},
		{
			name: "ERROR (NOT_RIGHT_TYPE)",
			input: struct {
				Object    interface{}
				fieldName string
				newValue  interface{}
			}{
				Object:    Data2{},
				fieldName: "ID",
				newValue:  "111",
			},
			expected: Data2{},
			err:      errors.New("obj должен быть структурой"),
		},
		{
			name: "ERROR (TYPE_MISMATCH)",
			input: struct {
				Object    interface{}
				fieldName string
				newValue  interface{}
			}{
				Object: Data2{
					ID:  "1",
					Arr: []int{1, 1, 1},
					Map: map[string]interface{}{
						"map": "map",
					},
					Any: "any",
				},
				fieldName: "ID",
				newValue:  1,
			},
			expected: Data2{},
			err:      errors.New("типы не совпадают: поле ID имеет тип string, а значение имеет тип int"),
		},
	}
	//	 Act
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			const caseOk = "OK"
			const caseNotFound = "ERROR (NOT_FOUND)"
			const caseNotRightType = "ERROR (NOT_RIGHT_TYPE)"
			const caseTypeMismatch = "ERROR (TYPE_MISMATCH)"

			switch test.name {
			case caseOk:
				data2, ok := test.input.Object.(Data2)

				err := ChangeField(&data2, test.input.fieldName, test.input.newValue)

				// Assert
				assert.Equal(t, true, ok)
				assert.NoError(t, err)
				assert.Equal(t, err, test.err)
				assert.Equal(t, test.expected, data2)
			case caseNotFound:
				data2, ok := test.input.Object.(Data2)

				err := ChangeField(&data2, test.input.fieldName, test.input.newValue)
				// Assert
				assert.Equal(t, true, ok)
				assert.Error(t, err)
				assert.Equal(t, err, test.err)
			case caseNotRightType:
				err := ChangeField(&test.input.Object, test.input.fieldName, test.input.newValue)

				// Assert
				assert.Error(t, err)
				assert.Equal(t, err, test.err)
			case caseTypeMismatch:
				data2, ok := test.input.Object.(Data2)

				err := ChangeField(&data2, test.input.fieldName, test.input.newValue)

				// Assert
				assert.Equal(t, true, ok)
				assert.Error(t, err)
				assert.Equal(t, err, test.err)
			}
		})
	}
}

func BenchmarkChangeFieldData2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data2 := Data2{ID: "1", Arr: []int{1, 1, 1}, Map: map[string]interface{}{"test": "test"}, Any: "any"}
		err := ChangeField(&data2, "ID", "2")
		require.NoError(b, err)
	}
}
