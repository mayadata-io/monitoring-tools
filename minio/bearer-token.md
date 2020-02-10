## Commands to get the Minio bearer-token for prometheus (for linux-amd64):

1. Download the `mc` cli
`wget https://dl.min.io/client/mc/release/linux-amd64/mc`
`chmod +x mc`

2. Add host to your application
mc config host add <ALIAS> <YOUR-MINIO-ENDPOINT> <YOUR-ACCESS-KEY> <YOUR-SECRET-KEY>
Example: `./mc config host add test-minio http://localhost:9000 minio minio111`

3. Generate Prometheus config for the alias
mc admin prometheus generate <ALIAS>
Example: `./mc admin prometheus generate test-minio`

From the above command's output, the bearer token will be present with the key as - `bearer_token`
