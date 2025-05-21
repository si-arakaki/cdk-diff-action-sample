# cdk-diff-action-sample

<img width="640" alt="cdk diff" src="https://github.com/user-attachments/assets/52c88ada-f035-48ee-a699-64e30382ee8a" />

## arn:aws:iam::123456789012:role/cdk-diff-action-sample

#### Trust relationships

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "Federated": "arn:aws:iam::123456789012:oidc-provider/token.actions.githubusercontent.com"
            },
            "Action": "sts:AssumeRoleWithWebIdentity",
            "Condition": {
                "StringEquals": {
                    "token.actions.githubusercontent.com:aud": "sts.amazonaws.com"
                },
                "StringLike": {
                    "token.actions.githubusercontent.com:sub": "repo:si-arakaki/*"
                }
            }
        }
    ]
}
```

#### Permissions

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "sts:AssumeRole",
            "Resource": "arn:aws:iam::123456789012:role/cdk-xxxxxxxxx-lookup-role-123456789012-ap-northeast-1"
        }
    ]
}
```
