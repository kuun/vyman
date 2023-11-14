#!/bin/bash

export VYMAN_USER=admin VYMAN_PASSWORD=password HTTP_API_KEY=vyman

sudo mkdir -p -m 0777 /var/run/vyman

./vyman --addr 0.0.0.0:8000

