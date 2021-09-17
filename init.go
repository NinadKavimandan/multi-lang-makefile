package main

/*
import (
	"encoding/json"
	"fmt"

	"cel/cel/cel1.com/example"

	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/interpreter/functions"
)

type Config struct {
	Name             string `yaml:"name"`
	WorkingDirectory string `yaml:"workingDirectory"`
	Settings         []struct {
		Name  string `yaml:"name"`
		Value string `yaml:"value"`
	} `yaml:"settings"`
}

type JsonFile struct {
	Dependencies struct {
		React string `json:"react"`
	} `json:"dependencies"`
}

func are_same(first ref.Val, second ref.Val) ref.Val {
	fmt.Println("here")
	if first == second {
		return types.True
	}
	return types.False
}

func init_cel_eval(exp string, contents string) ref.Val {
	var result JsonFile

	_ = json.Unmarshal([]byte(contents), &result)

	prot_json := &example.JsonFile{}

	prot_json.Dependencies = &example.DependenciesArray{}
	prot_json.Dependencies.React = result.Dependencies.React

	jsonFile := cel.Declarations(
		decls.NewVar("jsonFile", decls.NewObjectType("cel1.com.JsonFile")),
		decls.NewFunction("are_same", decls.NewOverload("are_same_string_string", []*expr.Type{decls.String, decls.String}, decls.Bool)),
		decls.NewFunction("are_not_same", decls.NewOverload("are_not_same_string_string", []*expr.Type{decls.String, decls.String}, decls.Bool)))
	env, _ := cel.NewEnv(cel.Types(&example.JsonFile{}), jsonFile)

	ast, iss := env.Compile(exp)
	if iss.Err() != nil {
		panic(iss.Err())
	}

	fmt.Println("Wohoooo! compiled!")

	prg, _ := env.Program(ast, cel.Functions(
		&functions.Overload{
			Operator: "are_same_string_string",
			Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
				return are_same(lhs, rhs)
			},
		},
		&functions.Overload{
			Operator: "are_not_same_string_string",
			Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
				fmt.Println("here")
				if lhs == rhs {
					return types.False
				}
				return types.True
			}},
	))
	out, _, _ := prg.Eval(map[string]interface{}{
		"jsonFile": prot_json,
	})

	fmt.Println(out)
	return out
}

/*
func main() {

	filename, _ := filepath.Abs("./config.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	json_filename, _ := filepath.Abs("./abc.json")
	jsonFilePlain, err2 := ioutil.ReadFile(json_filename)

	if err2 != nil {
		panic(err2)
	}

	var result JsonFile

	_ = json.Unmarshal(jsonFilePlain, &result)

	prot_json := &example.JsonFile{}

	prot_json.Dependencies = &example.DependenciesArray{}
	prot_json.Dependencies.React = result.Dependencies.React

	jsonFile := cel.Declarations(
		decls.NewVar("jsonFile", decls.NewObjectType("cel1.com.JsonFile")),
		decls.NewFunction("are_same", decls.NewOverload("are_same_string_string", []*expr.Type{decls.String, decls.String}, decls.Bool)),
		decls.NewFunction("are_not_same", decls.NewOverload("are_not_same_string_string", []*expr.Type{decls.String, decls.String}, decls.Bool)))
	env, _ := cel.NewEnv(cel.Types(&example.JsonFile{}), jsonFile)

	var config Config

	err1 := yaml.Unmarshal(yamlFile, &config)

	if err1 != nil {
		panic(err1)
	}

	fmt.Println(config.Name)

	fmt.Println("Settings are: ")

	for _, value := range config.Settings {
		fmt.Println(value.Name)

		if 1 == 1 {
			fmt.Println(value.Value)

			ast, iss := env.Compile(value.Value)
			if iss.Err() != nil {
				panic(iss.Err())
			}

			fmt.Println("Wohoooo! compiled!")

			prg, _ := env.Program(ast, cel.Functions(
				&functions.Overload{
					Operator: "are_same_string_string",
					Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
						return are_same(lhs, rhs)
					},
				},
				&functions.Overload{
					Operator: "are_not_same_string_string",
					Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
						fmt.Println("here")
						if lhs == rhs {
							return types.False
						}
						return types.True
					}},
			))
			out, _, _ := prg.Eval(map[string]interface{}{
				"jsonFile": prot_json,
			})

			fmt.Println(out)
			fmt.Println(out.Type())
			res := bool(out.(types.Bool))
			fmt.Println(res)
		}
	}
}
*/
