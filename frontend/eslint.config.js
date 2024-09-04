import path from "node:path";
import { fileURLToPath } from "node:url";

import { fixupPluginRules } from "@eslint/compat";
import { FlatCompat } from "@eslint/eslintrc";
import js from "@eslint/js";
import typescriptEslint from "@typescript-eslint/eslint-plugin";
import tsParser from "@typescript-eslint/parser";
import eslintPluginImportX from "eslint-plugin-import-x";
import react from "eslint-plugin-react";
import reactHooksPlugin from "eslint-plugin-react-hooks";
import globals from "globals";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);
const compat = new FlatCompat({
  baseDirectory: __dirname,
  recommendedConfig: js.configs.recommended,
  allConfig: js.configs.all,
});

export default [
  {
    ignores: ["**/node_modules", "**/build/", "src/vendor/*.js"],
  },
  ...compat.extends(
    "eslint:recommended",
    "plugin:react/jsx-runtime",
    "plugin:react/recommended",
    "plugin:@typescript-eslint/recommended",
    "plugin:prettier/recommended"
  ),
  {
    plugins: {
      react,
      "@typescript-eslint": typescriptEslint,
      "import-x": eslintPluginImportX,
      "react-hooks": fixupPluginRules(reactHooksPlugin),
    },

    languageOptions: {
      globals: {
        ...globals.browser,
        ...globals.node,
      },

      parser: tsParser,
      ecmaVersion: "latest",
      sourceType: "module",
    },

    settings: {
      ...eslintPluginImportX.flatConfigs.typescript.settings,
    },

    rules: {
      ...reactHooksPlugin.configs.recommended.rules,
      "import-x/order": [
        "error",
        {
          groups: ["builtin", "external", "internal", "parent", "sibling", "index", "object"],

          pathGroups: [
            {
              pattern: "react",
              group: "external",
              position: "before",
            },
            {
              pattern: "**/*.css",
              group: "object",

              patternOptions: {
                matchBase: true,
              },

              position: "after",
            },
          ],

          pathGroupsExcludedImportTypes: ["react"],
          "newlines-between": "always",

          alphabetize: {
            order: "asc",
            caseInsensitive: true,
          },
        },
      ],

      "no-unused-vars": "off",
      "@typescript-eslint/no-unused-vars": "error",
      "import-x/no-unresolved": "error",
      "import-x/named": "error",
      "import-x/namespace": "error",
      "import-x/default": "error",
      "import-x/export": "error",
    },
  },
];
