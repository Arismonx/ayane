# Ayane

```
          
              ___                            
             /   |  __  __ ____ _ ____   ___ 
            / /| | / / / // __  // __ \ / _ \
           / ___ |/ /_/ // /_/ // / / //  __/
          /_/  |_|\__, / \__,_//_/ /_/ \___/ 
                 /____/                      
          
          => Ayane CLI Todo Application v1.0.0 


```

**Ayane** is a simple and intuitive Command Line Interface (CLI) application for managing your daily tasks. It allows you to quickly `add`, `list`, `mark` as `complete`, `edit`, and `delete` your to-do tesk directly from your terminal.

## Features
- **Add Tasks:** Quickly add new tasks to your list.
- **List Tasks:** View all your tasks, showing their status (completed/incomplete).
- **Mark Tasks as Complete:** Mark tasks as done, including a completion timestamp.
- **Edit Tasks:** Modify the description of existing tasks.
- **Delete Tasks:** Remove tasks from your list.
- **Persistent Storage:** All tasks are saved to a local `JSON` file, so your list persists across sessions.
## Package used
Used for creating [table](https://github.com/aquasecurity/table)
## Installation
To install the Ayane CLI, you need to have [Go](https://go.dev/doc/install) (version 1.18 or higher) installed on your system. 
> [download](https://github.com/Arismonx/ayane/releases/tag/v1.0.0) Ayane CLI v1.0.0

## Usage
Suppose you downloaded the `ayane` file and stored it in your Downloads folder.

- **Windows:**
Locate the file: Go to the folder where you downloaded the file. You should see a file named `ayane_windows_amd64.exe`

>Run command in `cmd` or `PowerShell`
```
ayane_windows_amd64.exe
```

- **Linux / macOS:**
>Run command
```
cd ~/Downloads
```
You should see a file named `ayane_linux_amd64` for Linux or `ayane_darwin_amd64` for MacOS

>Give permission to run
 - for Linux
```
chmod +x ayane_linux_amd64
```
and Run command for Linux
```
./ayane_linux_amd64
```
- for MacOS
```
chmod +x ayane_darwin_amd64
```
and Run command for MacOS
```
./ayane_darwin_amd64
```
 **After run command you should see**
```
              ___                            
             /   |  __  __ ____ _ ____   ___ 
            / /| | / / / // __  // __ \ / _ \
           / ___ |/ /_/ // /_/ // / / //  __/
          /_/  |_|\__, / \__,_//_/ /_/ \___/ 
                 /____/                      
          
          => Ayane CLI Todo Application v1.0.0 
```
> The program is ready to work. :white_check_mark:

> [!IMPORTANT]
> Don't forget to read the [user manual]()  >W<.

## User manual
Here are the commands you can use.
```
All Command: 
  -add       Add a new todo by specify task
  -edit      Edit a tody by index & specify task
  -del       Specify a todo by index to delete
  -list      List all todo
  -mark      Specify a todo by index to mark
  -help      All Command
  -ayane     home page
```
- How to use command:
  
  - command `-add` use when you want to `add` a `task`.
    
  ```
  ./ayane_linux_amd64 -add "Go Run!"
  ./ayane_linux_amd64 -add "Do homework"
  ```
  
  - command `-edit` use when you want to `edit` a task by use `id` of task.


  ```
  ./ayane_linux_amd64 -edit 0:"Go Exercise!"
  ```


  - command `-del` use when you want to `delete` a task by use `id` of task.
         
  ```
  ./ayane_linux_amd64 -del 0
  ```
  
  - command `-list` use when you want view a list of all tasks.


  ```
  ./ayane_linux_amd64 -list
  ```

  - command `-mark` Use when you have completed the task by use id of task.

  ```
  ./ayane_linux_amd64 -mark 0
  ```

  - command `-help` Use when you want see command.

  ```
  ./ayane_linux_amd64 -help
  ```

  - command `-ayane` nothing just cute >_< .
  ```
  ./ayane_linux_amd64 -ayane
  ```

**Output Example**
```go
┌────┬─────────────┬───────┬──────────────────────────┬──────────────────────────┐
│ ID │    Task     │ Check │        CreatedAt         │        CompleteAt        │
├────┼─────────────┼───────┼──────────────────────────┼──────────────────────────┤
│ 0  │ Do homework │ [✓]   │ Thu, 26 Jun 2025, 2:32PM │ Thu, 26 Jun 2025, 2:33PM │
└────┴─────────────┴───────┴──────────────────────────┴──────────────────────────┘
```
**And you will get a `todos.json` file.**
```json
[
  {
    "Task": "Do homework",
    "Check": true,
    "CreatedAt": "Thu, 26 Jun 2025, 2:32PM",
    "CompleteAt": "Thu, 26 Jun 2025, 2:33PM"
  }
]
```
## Contributing
Contributions are welcome! If you find a bug or have a feature request, please open an issue on the GitHub repository. If you'd like to contribute code, feel free to fork the repository and submit a pull request.
  

  
