package extract

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	. "github.com/TykTechnologies/exp/cmd/schema-gen/model"
)

func TestGetDeclarations(t *testing.T) {
	for _, tt := range getTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			tmpFile := filepath.Join(tmpDir, "test.go")
			err := os.WriteFile(tmpFile, []byte(tt.code), 0o644)
			if err != nil {
				t.Fatal(err)
			}

			pkgInfos, err := Extract(tmpFile, tt.options)
			if err != nil {
				t.Fatal(err)
			}

			if len(pkgInfos) != 1 {
				t.Fatalf("expected 1 package, got %d", len(pkgInfos))
			}

			got := pkgInfos[0]
			gotJSON, _ := json.MarshalIndent(got, "", "    ")
			expectedJSON, _ := json.MarshalIndent(tt.expected, "", "    ")

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("\nGot:\n%s\n\nExpected:\n%s", string(gotJSON), string(expectedJSON))
			}
		})
	}
}

func TestGetDeclarationsWithFunctions(t *testing.T) {
	code := `package test
    type User struct {
        ID   int
        Name string
    }
    
    func (u *User) GetName() string {
        return u.Name
    }
    
    func GlobalFunc() int {
        return 42
    }`

	options := &ExtractOptions{includeFunctions: true}

	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.go")
	err := os.WriteFile(tmpFile, []byte(code), 0o644)
	if err != nil {
		t.Fatal(err)
	}

	pkgInfos, err := Extract(tmpFile, options)
	if err != nil {
		t.Fatal(err)
	}

	if len(pkgInfos[0].Functions) != 1 {
		t.Errorf("expected 1 global function, got %d", len(pkgInfos[0].Functions))
	}

	userType := pkgInfos[0].Declarations[0].Types[0]
	if len(userType.Functions) != 1 {
		t.Errorf("expected 1 method for User, got %d", len(userType.Functions))
	}

	if userType.Functions[0].Name != "GetName" {
		t.Errorf("expected method name GetName, got %s", userType.Functions[0].Name)
	}
}

func getTestCases() []struct {
	name     string
	code     string
	options  *ExtractOptions
	expected *PackageInfo
} {
	return []struct {
		name     string
		code     string
		options  *ExtractOptions
		expected *PackageInfo
	}{
		{
			name: "Basic type with int enum",
			code: `package test
type Status int
const (
	Active Status = iota
	Inactive
	Suspended
)`,
			options: &ExtractOptions{},
			expected: &PackageInfo{
				Imports: []string{},
				Name:    "test",
				Declarations: DeclarationList{
					{
						Types: TypeList{
							{
								Name: "Status",
								Type: "int",
								EnumValues: []*EnumValue{
									{Name: "Active", Value: 0},
									{Name: "Inactive", Value: 1},
									{Name: "Suspended", Value: 2},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "String enum",
			code: `package test
type Role string
const (
	Admin Role = "admin"
	User Role = "user"
	Guest Role = "guest"
)`,
			options: &ExtractOptions{},
			expected: &PackageInfo{
				Imports: []string{},
				Name:    "test",
				Declarations: DeclarationList{
					{
						Types: TypeList{
							{
								Name: "Role",
								Type: "string",
								EnumValues: []*EnumValue{
									{Name: "Admin", Value: "admin"},
									{Name: "User", Value: "user"},
									{Name: "Guest", Value: "guest"},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Multiple types and enums",
			code: `package test
type Status int
type Role string
const (
	Active Status = iota
	Inactive
	Suspended
)
const (
	Admin Role = "admin"
	User Role = "user"
)`,
			options: &ExtractOptions{},
			expected: &PackageInfo{
				Imports: []string{},
				Name:    "test",
				Declarations: DeclarationList{
					{
						Types: TypeList{
							{
								Name: "Status",
								Type: "int",
								EnumValues: []*EnumValue{
									{Name: "Active", Value: 0},
									{Name: "Inactive", Value: 1},
									{Name: "Suspended", Value: 2},
								},
							},
						},
					},
					{
						Types: TypeList{
							{
								Name: "Role",
								Type: "string",
								EnumValues: []*EnumValue{
									{Name: "Admin", Value: "admin"},
									{Name: "User", Value: "user"},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Unexported Types",
			code: `package test
type status int
const (
	active status = iota
	inactive
)`,
			options: &ExtractOptions{includeUnexported: false},
			expected: &PackageInfo{
				Imports:      []string{},
				Name:         "test",
				Declarations: DeclarationList{},
			},
		},
		{
			name: "Type with functions",
			code: `package test
type User struct {
	ID   int
	Name string
}
func (u *User) GetName() string {
	return u.Name
}`,
			options: &ExtractOptions{includeFunctions: true},
			expected: &PackageInfo{
				Imports: []string{},
				Name:    "test",
				Declarations: DeclarationList{
					{
						Types: TypeList{
							{
								Name: "User",
								Fields: []*FieldInfo{
									{Name: "ID", Type: "int", Path: "User.ID", JSONName: "ID"},
									{Name: "Name", Type: "string", Path: "User.Name", JSONName: "Name"},
								},
								Functions: []*FuncInfo{
									{
										Name:      "GetName",
										Path:      "User",
										Type:      "u *User",
										Signature: "GetName () string",
										Source:    "func (u *User) GetName() string {\n\treturn u.Name\n}",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
