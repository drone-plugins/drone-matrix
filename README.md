# drone-plugin-matrix

[Drone](https://drone.io/) notifications to [Matrix](https://matrix.org/)

Usage:

```yaml
  matrix:
    image: ptman/drone-plugin-matrix
    homeserver: https://matrix.org          # defaults to https://matrix.org
    roomid: '!0123456789abcdef:matrix.org'  # room has to already be joined
    secrets:
      - matrix_username     # either username  ('ourbot')
      - matrix_password     # and password     ('*ourbot-password*')
    # - matrix_userid       # or userid        ('@ourbot:matrix.org')
    # - matrix_accesstoken  # and access token ('long string of characters')
```

## License

ISC
