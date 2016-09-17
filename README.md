# Homenum Revelio

An idea about displaying who is currently at home, thanks to [Dan Slocombe](https://github.com/danslocombe).

*This project is very much experimental and a bit of fun.*

## Dependencies

 - `nmap`, though an earlier commit [c269890](https://github.com/egnwd/homenum-revelio/commits/c2698904ba66014c3e3aef23b152c0aac56343f3) uses `ping`

## Running
 - You will need a yaml file mapping MAC addresses to Names in the form:

```yaml
05:E5:87:F4:4C:D4: Name
59:65:11:03:56:68: Other Name
38:F6:AC:4D:ED:09: Full House
```
 - Then run using `./who_is_home` from within the directory or `who_is_home` by adding the folder to your `PATH`
