# aws-env
Export your specified profile credentials to environment variable

## Usage

Specify profile using flag

```shell
$ aws-env --profile YOUR_AWS_PROFILE

```

or just `aws-env` it will use `default` profile or the profile you define in configuration file
`$HOME/.aws-env/config.yaml`

## Installation
Download the latest release binary [here](https://github.com/Gujarats/aws-env/releases) and extract to your `bin` for example `/usr/local/bin/`.
 
 * Create file `$HOME/.aws-env/config.yaml`
 * specify your profile

 ```yaml
 profile: your-custome-profile
 ```
