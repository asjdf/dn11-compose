#!/bin/bash

set -e
set -x

# 初始化所有基于wg的peers连接
if ls /etc/wireguard/*.conf ; then
    for i in /etc/wireguard/*.conf; do wg-quick up $i; done
fi

# 启动bird
bird -d