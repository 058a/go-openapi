docker run --rm -v ./:/sqlboiler --network deployments_default curvegrid/sqlboiler:psql psql --output ./sqlboiler --pkgname sqlboiler && sudo chown -R k8suser ./sqlboiler && mv ./sqlboiler ./../../internal/infra/ && go mod tidy