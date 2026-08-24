import boto3


class R2Storage:

    def __init__(
        self,
        account_id: str,
        access_key_id: str,
        secret_access_key: str,
        bucket: str,
    ):
        self.bucket = bucket

        self.client = boto3.client(
            "s3",
            endpoint_url=(
                f"https://{account_id}.r2.cloudflarestorage.com"
            ),
            aws_access_key_id=access_key_id,
            aws_secret_access_key=secret_access_key,
            region_name="auto",
        )

    def download(
        self,
        storage_key: str,
    ) -> bytes:

        response = self.client.get_object(
            Bucket=self.bucket,
            Key=storage_key,
        )

        return response["Body"].read()