# drone-plugin-matrix

[Drone](https://drone.io/) notifications to [Matrix](https://matrix.org/)

Usage:

```yaml
  matrix:
    image: ptman/drone-plugin-matrix
    homeserver: https://matrix.org
    roomid: '!0123456789abcdef:matrix.org'  # room has to already be joined
    username: ourbot                        # either username
    password: *account-password*            # and password
    userid: @ourbot:matrix.org              # or userid
    accesstoken: 0123456789abcdef           # and accesstoken
    secrets:                                # and a better idea
      - source: matrix_username             # is to not store
        target: plugin_username             # credentials in the git repo
      - source: matrix_password             # but instead use drone
        target: plugin_password             # secret management
```

## License

ISC
