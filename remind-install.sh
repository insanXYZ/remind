wget https://github.com/insanXYZ/remind/releases/download/v1.0.0/remind-daemon
wget https://github.com/insanXYZ/remind/releases/download/v1.0.0/remind

sudo chmod a+x ./remind
sudo chmod a+x ./remind-daemon

sudo mv ./remind /usr/local/bin
sudo mv ./remind-daemon /usr/local/bin

SERVICE_NAME="remind-daemon"
EXEC_PATH="/usr/local/bin/$SERVICE_NAME"
USER=$(whoami)

# create systemd file
cat <<EOF | sudo tee /etc/systemd/system/$SERVICE_NAME.service
[Unit]
Description=Remind Daemon Service
After=network.target

[Service]
ExecStart=$EXEC_PATH
Restart=always
User=$USER
Environment="DISPLAY=:0"
Environment="DBUS_SESSION_BUS_ADDRESS=unix:path=/run/user/1000/bus"

[Install]
WantedBy=multi-user.target
EOF

# Reload systemd dan enable service
sudo systemctl daemon-reload
sudo systemctl enable $SERVICE_NAME.service
sudo systemctl start $SERVICE_NAME.service