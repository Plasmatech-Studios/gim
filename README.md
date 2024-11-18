## GIM: Git Identity Manager

---

**GIM** (Git Identity Manager) is a CLI tool designed to help developers manage multiple SSH keys effortlessly. Whether you're juggling work and personal GitHub accounts or multiple project identities, GIM simplifies the process of switching between them.

---

### Features

- **Add SSH Keys**: Generate new SSH keys and configure them under an alias.
- **List SSH Keys**: View all configured SSH keys, including orphaned keys.
- **Switch Keys**: Seamlessly switch between different SSH keys.
- **Remove Keys**: Remove aliases and optionally delete associated SSH key files.
- **Restore Orphaned Keys**: Re-add keys that are present in your system but not configured.
- **Rename Aliases**: Rename existing aliases without affecting the underlying keys.
- **Track Active Key**: Easily see which SSH key is currently in use.
- **Dynamic Help**: Get detailed usage information for each command.

---

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/your-username/gim.git
   cd gim
   ```

2. **Build the executable**:

   ```bash
   go build -o gim cmd/gim.go
   ```

3. **Add `gim` to your PATH** (optional for global access):
   ```bash
   mv gim /usr/local/bin/
   ```

---

### Usage

GIM provides the following commands:

| Command   | Description                                                             |
| --------- | ----------------------------------------------------------------------- |
| `list`    | Lists all configured SSH keys. Use `-a` to include orphaned keys.       |
| `add`     | Adds a new SSH key with a given alias.                                  |
| `use`     | Switches to the specified SSH key.                                      |
| `remove`  | Removes a key alias. Use `-d` to delete the key files.                  |
| `restore` | Restores an orphaned SSH key under a given alias.                       |
| `rename`  | Renames an existing alias.                                              |
| `using`   | Displays the currently active SSH key. Use `-c` to copy its public key. |
| `help`    | Displays usage information for all commands.                            |

---

### Examples

**Add a new key**:

```bash
gim add work
```

**List keys**:

```bash
gim list
gim list -a  # Include orphaned keys
```

**Use a key**:

```bash
gim use work
```

**Remove an alias and delete its files**:

```bash
gim remove -d personal
```

**Restore an orphaned key**:

```bash
gim restore old_key
```

**Rename an alias**:

```bash
gim rename work work_backup
```

**Show the active key**:

```bash
gim using
```

**Show the active key and copy its public key to the clipboard**:

```bash
gim using -c
```

---

### Contribution

Contributions are welcome! If you have ideas or improvements, feel free to open an issue or submit a pull request.

---

### License

This project is licensed under the **Pete's Non-Commercial License**.

#### Key Terms:

- **Non-Commercial Use Only**: You may use, modify, and share this software, but **commercial use is prohibited** without explicit written permission.
- **No AI Training**: This software **cannot be used to train or improve AI or machine learning models** without explicit written permission.
- **Attribution Required**: If you share or modify the software, you must include the original license and attribution.

For full details, see the [LICENSE](LICENSE) file.

---
