package config

import (
	"fmt"
	"reflect"
	"strings"
)

// FieldDoc holds documentation for a config field
type FieldDoc struct {
	Name        string
	Type        string
	Default     string
	Example     string
	Description string
	Validation  string
	Path        string // e.g., "app.name"
}

// GetConfigDocs returns documentation for all config fields
func GetConfigDocs() []FieldDoc {
	var docs []FieldDoc
	collectDocs("", reflect.TypeOf(Config{}), &docs)
	return docs
}

// collectDocs recursively collects documentation from struct fields
func collectDocs(prefix string, typ reflect.Type, docs *[]FieldDoc) {
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		// Get mapstructure tag for the key name
		mapstructureTag := field.Tag.Get("mapstructure")
		if mapstructureTag == "" || mapstructureTag == "-" {
			continue
		}

		// Build the full key path
		var path string
		if prefix != "" {
			path = prefix + "." + mapstructureTag
		} else {
			path = mapstructureTag
		}

		// Handle nested structs
		if field.Type.Kind() == reflect.Struct {
			collectDocs(path, field.Type, docs)
			continue
		}

		// Extract tags
		doc := FieldDoc{
			Name:        field.Name,
			Type:        field.Type.String(),
			Default:     field.Tag.Get("default"),
			Example:     field.Tag.Get("example"),
			Description: field.Tag.Get("doc"),
			Validation:  field.Tag.Get("validate"),
			Path:        path,
		}

		*docs = append(*docs, doc)
	}
}

// GenerateYAMLExample generates an example config.yaml with comments
func GenerateYAMLExample() string {
	var sb strings.Builder
	docs := GetConfigDocs()

	// Group by top-level section
	sections := make(map[string][]FieldDoc)
	for _, doc := range docs {
		parts := strings.SplitN(doc.Path, ".", 2)
		section := parts[0]
		sections[section] = append(sections[section], doc)
	}

	// Generate YAML
	sb.WriteString("# Application Configuration\n")
	sb.WriteString("# Auto-generated from struct tags\n\n")

	for section, fields := range sections {
		sb.WriteString(fmt.Sprintf("%s:\n", section))
		
		for _, field := range fields {
			// Add documentation comment
			if field.Description != "" {
				sb.WriteString(fmt.Sprintf("  # %s\n", field.Description))
			}
			
			// Add validation info
			if field.Validation != "" {
				sb.WriteString(fmt.Sprintf("  # Validation: %s\n", field.Validation))
			}
			
			// Add example
			if field.Example != "" {
				sb.WriteString(fmt.Sprintf("  # Example: %s\n", field.Example))
			}
			
			// Add default value comment
			if field.Default != "" {
				sb.WriteString(fmt.Sprintf("  # Default: %s\n", field.Default))
			}
			
			// Get field name (remove section prefix)
			parts := strings.SplitN(field.Path, ".", 2)
			var fieldName string
			if len(parts) == 2 {
				fieldName = parts[1]
			} else {
				// Top-level field without section
				fieldName = parts[0]
			}
			
			// Add the actual field with default value
			value := field.Default
			if value == "" {
				value = field.Example
			}
			
			// Quote strings (check if base type is string)
			if strings.Contains(field.Type, "string") && value != "" {
				value = fmt.Sprintf(`"%s"`, value)
			}
			
			sb.WriteString(fmt.Sprintf("  %s: %s\n", fieldName, value))
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

// GenerateMarkdownDocs generates markdown documentation table
func GenerateMarkdownDocs() string {
	var sb strings.Builder
	docs := GetConfigDocs()

	sb.WriteString("# Configuration Reference\n\n")
	sb.WriteString("| Field | Type | Default | Example | Description | Validation |\n")
	sb.WriteString("|-------|------|---------|---------|-------------|------------|\n")

	for _, doc := range docs {
		sb.WriteString(fmt.Sprintf("| `%s` | %s | `%s` | `%s` | %s | %s |\n",
			doc.Path,
			doc.Type,
			doc.Default,
			doc.Example,
			doc.Description,
			doc.Validation,
		))
	}

	return sb.String()
}

// PrintConfigHelp prints human-readable config documentation
func PrintConfigHelp() {
	docs := GetConfigDocs()
	
	fmt.Println("Configuration Fields:")
	fmt.Println(strings.Repeat("=", 80))
	
	for _, doc := range docs {
		fmt.Printf("\n%s\n", doc.Path)
		fmt.Println(strings.Repeat("-", len(doc.Path)))
		
		if doc.Description != "" {
			fmt.Printf("Description: %s\n", doc.Description)
		}
		
		fmt.Printf("Type: %s\n", doc.Type)
		
		if doc.Default != "" {
			fmt.Printf("Default: %s\n", doc.Default)
		}
		
		if doc.Example != "" {
			fmt.Printf("Example: %s\n", doc.Example)
		}
		
		if doc.Validation != "" {
			fmt.Printf("Validation: %s\n", doc.Validation)
		}
	}
}
