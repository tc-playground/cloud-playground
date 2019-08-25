# AWS CLI

## Introduction

---

## Install

* [AWS CLI](https://aws.amazon.com/cli/)

* [Install AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/install-linux.html)

* [Troubleshoot AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-troubleshooting.html)

---

## Configure

* Create `.aws` configuration via the `aws` CLI:

    ```
    aws configure
    ```

* `$HOME/.aws/config` contains a default `ini` configuration file.

    ```
    [default]
    output = text
    region = eu-west-2
    ```

* `$HOME/.aws/credentials` contains a default `ini` configuration file.

    ```
    [default]
    aws_access_key_id = XXX 
    aws_secret_access_key = YYY
    
    [trjl-dev]
    aws_access_key_id = III
    aws_secret_access_key = JJJ
    ```

---

## AWS Access Keys

* Access keys consist of two parts: 

    * An access key `ID` - e.g. AKIAIOSFODNN7EXAMPLE
    
    * An access key `secret` - e.g. wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY

* Create an `access key` in `iam` for a `user`.

* Store via `environment variables` or in `./aws/credentials` file.

---

## Example Usage

* __Man page__ : `aws help`

* __List Users__ :  `aws iam list-users`

* __List Instances__ : `aws ec2 describe-instances`