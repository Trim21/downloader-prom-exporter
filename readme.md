A prometheus exporter for rtorrent, transmissionBT or qbittorrent.

```yaml
version: "3.7"

service:
  downloader-prom-exporter:
    image: ghcr.io/trim21/downloader-prom-exporter:latest
    restart: unless-stopped
    environment:
      RTORRENT_API_ENTRYPOINT: "http://rtorrent.omv.trim21.me/RPC2"
      # or RTORRENT_API_ENTRYPOINT: "scgi://127.0.0.1:5000"
      # don't forget to mount socket file into container
      # or RTORRENT_API_ENTRYPOINT: "scgi:////home/ubuntu/.local/share/.rtorrent.sock"
      TRANSMISSION_API_ENTRYPOINT: "http://admin:password@192.168.1.3:8080"
      QBIT_API_ENTRYPOINT: "https://qb.omv.trim21.me"
    ports:
      - "8521:80"
```

In your prometheus scrape config

```yaml
scrape_configs:
  - job_name: "downloader-exporter"
    static_configs:
      - targets: ["127.0.0.1:8521"]
```

## dashboard

https://grafana.com/grafana/dashboards/18986


## Tips

Some downloaders' rpc thread may freeze in some situation, this will block downloader exporter.

If you want to avoid this, please run each exporter instance for each downloader.
