# DN11-Compose

## Tool

- Prepare privKey and pubKey for WireGuard

  ```sh
  wg genkey | tee /dev/tty | wg pubkey
  ```

## Setup

We recommend to use [docker-compose v2](https://docs.docker.com/compose/cli-command/#install-on-linux) to setup containers. 

- Build all containers and docker networks
  
  ```sh
  docker compose up --build --no-start
  ```

- Start all containers

  ```sh
  docker compose start
  ```
  
- Show status of containers

  ```sh
  docker compose ps
  ```

- Enter specific container

  ```sh
  docker compose exec bgp bash
  ```

- Stop and delete all services and docker networks completely

  ```sh
  docker compose down
  ```

## Config

> We only list the minimum configuration needed to make it work. For more configurations, please read the code or [docker-compose specification](https://docs.docker.com/compose/compose-file/)


1. Configure the `subnet` for `dn11-net` in `docker-compose.yml`

    ```yml
    networks:
      dn11-net:
        driver: bridge
        enable_ipv6: true
        internal: false
        ipam:
          driver: default
          config:
            - subnet: <your dn11 ipv4 subnet>
            - subnet: <your dn11 ipv6 subnet>
    ```

    > !!! Note that the docker host will take up **the first address on the subnet**. So you cannot assign the first ip to any of the containers.

2. For each service, you may want to assign a dn42 ip address fot it.

    ```yml
        networks:
          dn42-net:
              ipv4_address: "<dn11 ip address allocated this service>"
              ipv6_address: "<dn11 ip address allocated this service>"
    ```

3. All containers except the bgp container need to manually configure routes to forward traffic going to the dn42 network to the bgp container.

   We provide two environment variable to configure the ip address of the dn42 gateway

    ```yml
        environment:
          - DN42_GATEWAY_V4=<ipv4 address of your bgp container>
          - DN42_GATEWAY_V6=<ipv6 address of your bgp container>
    ```

## peer

1. Edit `peer/named.conf`

   You need to edit the config file of bird2 `bgp/named.conf` according to the guidance [here](https://dn42.eu/howto/Bird2).

2. For each of your peers, create files in dir `bgp/bird2-peers` and `bgp/wg-peers`

3. You need to add port mapping for your peers in `docker-compose.yml`

   ```yml
       ports:
        - "12345:12345/udp" # export for peer（wireguard）
   ```
