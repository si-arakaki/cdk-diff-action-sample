# cdk-diff-action-sample

> <img width="640" alt="cdk diff" src="https://github.com/user-attachments/assets/30ea67e8-c47b-4443-b4e8-719145650e0f" />

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
