<p align="center">
    <img src="https://github.com/user-attachments/assets/03067a06-abb5-445f-ba6a-0963af32e7fe" width="240"/>
</p>

# Remind - reminder tool on your linux

<video src="https://github.com/user-attachments/assets/918f7ee4-dd2a-4381-b605-da91a8988702"/>

# 💾 Installation
Before you install it, you must have the ``notify-send`` dependency. 

```bash
# Arch linux
sudo pacman -S libnotify

# Debian / Ubuntu
sudo apt install libnotify-bin

# Fedora
sudo dnf install libnotify
```

After that you can do this

```bash
wget https://github.com/insanXYZ/remind/releases/download/v1.0.0/remind-install.sh

sudo chmod a+x ./remind-install.sh

./remind-install.sh
```

# Usage
some examples of using remind
```bash
# create remind
remind set --name "go to school" --time "07:30"

# delete remind
remind delete --id 1

# check remind
remind check --id 1

# list all remind
remind ls
```

# License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.