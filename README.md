# IMPI Prometheus exporter. 

The idea behind this exporter is to expose the power measurements given by the following command:

```bash
ipmi-sensors -h localhost -u <bmc-user> -p <bmc-password> --sensor-types Current
```