# DDRaceTemplateEditor

Welcome to DDRaceTemplateEditor! This repository is a template editor for DDRace, a popular multiplayer mod for Teeworlds.

## Using go-ddrtemplateeditor

The go-ddrtemplateeditor is a simple CLI tool to exchange items from one template with an other.

### Prerequisites

1. Install [Go](https://go.dev/doc/install)
2. Clone the repository: `git clone https://github.com/your-username/DDRaceTemplateEditor.git`

### Usage

Since it is (for now) just a simple CLI tool, the usage is pretty easy. 
First, open a terminal navigate to the source folder and build the executable:

```bash
cd <your-path>/DDRaceTemplateEditor/cmd/go-ddrtemplateeditor
go build
```

Now run the tool like follows:

```bash
go-ddrtemplateeditor -item hammer,sword -src template1.png -dst template2.png -out output.png
```

The `-item` parameter defines, which items will be extracted from the source template (`src`) and pasted into the destination template (`-dst`). The result of the exchange will be saved as the file defined with the `-out` parameter.

To list all available items, just use `./go-ddrtemplateeditor -h`.

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or bug reports, please open an issue or submit a pull request.

## Todo
- Create more items
- Create an UI