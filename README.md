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
cd <your-path>/DDRaceTemplateEditor/go-ddrtemplateeditor/cmd/cli
go build -o go-ddrtemplateeditor
```

Now run the tool like follows:

```bash
./go-ddrtemplateeditor -item hammer,sword -src template1.png -dst template2.png -out output.png
```

The `-item` parameter defines, which items will be extracted from the source template (`src`) and pasted into the destination template (`-dst`). The result of the exchange will be saved as the file defined with the `-out` parameter.

To list all available items, just use `./go-ddrtemplateeditor -h`.

## Build the dockerfile

Build the dockerfile with the following command:

```bash
docker build --pull --rm -f "go-ddrtemplateeditor\dockerfile" -t ddracetemplateeditor:latest "go-ddrtemplateeditor"
```

Now you can run the docker container with the following command:

```bash
docker run --rm -it -p 1337:1337/tcp ddracetemplateeditor:latest 
```

The container will start a simple webserver on port 1337. You can access the webserver by opening your browser and navigating to `http://localhost:1337/api/templates`.

## Test with PowerShell

You can test the API with PowerShell. Here is an example:

```powershell
$uri = "http://localhost:1337/api/templates"
$response = Invoke-RestMethod -Uri $uri -Method GET 
$response
```

To download a template, you can use the following command:

```powershell
$uri = "http://localhost:1337/api/templates/1/image"
Invoke-WebRequest -Uri $uri -Method GET -OutFile "template1.png"
```

To exchange an item from one template to an other, you can use the following command:

```powershell
$uri = "http://localhost:1337/api/templates/1/replace"
$body = @{
    templateId = 2
    items = @(
        @{
            id = 1
        }
    )
} | ConvertTo-Json
Invoke-RestMethod -Uri $uri -Method PUT -Body $body -ContentType "application/json"
```

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or bug reports, please open an issue or submit a pull request.

## Todo
- Create more items
- Create an UI