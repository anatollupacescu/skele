# README

This is a basic syntax highlighter for the `skele` test specification language.

## Colours

Add this to your VSCode's `settings.json` file:

```json
    "editor.tokenColorCustomizations": {
        "textMateRules": [
            {
                "scope": "keyword.comment.skl",
                "settings": {
                    "foreground": "#829294"
                }
            },{
                "scope": "keyword.control.skl",
                "settings": {
                    "foreground": "#a7ab3f"
                }
                },{
                    "scope": "keyword.bold.skl",
                    "settings": {
                        "foreground": "#cd5424"
                    }
                }
        ]
    }
```
