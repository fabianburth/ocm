## ocm componentversions sign &mdash; Sign Component Version

### Synopsis

```
ocm componentversions sign [<options>] {<component-reference>}
```

### Options

```
      --algorithm string          signature handler
  -h, --help                      help for sign
  -K, --private-key stringArray   private key setting
  -k, --public-key stringArray    public key setting
  -r, --repo string               repository name or spec
  -s, --signature string          signature name
      --update                    update digest in component versions (default true)
  -V, --verify                    verify existing digests (default true)
```

### Description


Sign specified component versions. 

If the <code>--repo</code> option is specified, the given names are interpreted
relative to the specified repository using the syntax

<center>
    <pre>&lt;component>[:&lt;version>]</pre>
</center>

If no <code>--repo</code> option is specified the given names are interpreted 
as located OCM component version references:

<center>
    <pre>[&lt;repo type>::]&lt;host>[:&lt;port>][/&lt;base path>]//&lt;component>[:&lt;version>]</pre>
</center>

Additionally there is a variant to denote common transport archives
and general repository specifications

<center>
    <pre>[&lt;repo type>::]&lt;filepath>|&lt;spec json>[//&lt;component>[:&lt;version>]]</pre>
</center>

The <code>--repo</code> option takes an OCM repository specification:

<center>
    <pre>[&lt;repo type>::]&lt;configured name>|&lt;file path>|&lt;spec json></pre>
</center>

For the *Common Transport Format* the types <code>directory</code>,
<code>tar</code> or <code>tgz</code> is possible.

Using the JSON variant any repository type supported by the 
linked library can be used:

Dedicated OCM repository types:
- `ComponentArchive`

OCI Repository types (using standard component repository to OCI mapping):
- `ArtefactSet`
- `CommonTransportFormat`
- `DockerDaemon`
- `Empty`
- `OCIRegistry`
- `oci`

The <code>--public-key</code> and <code>--private-key</code> optiond can be
used to define public and provate keys on the command line. The options have an
argument of the form <code>[&lt;name>=]&lt;filepath></code>. The optional name
specified the signature name the key should be used for. By default this is the
signature name specified with the option <code>--signature</code>.


### Examples

```
$ ocm sign componentversion --signature mandelsoft --private-key=mandelsoft.key ghcr.io/mandelsoft/kubelink
```

### SEE ALSO

##### Parents

* [ocm componentversions](ocm_componentversions.md)	 - Commands acting on components
* [ocm](ocm.md)	 - Open Component Model command line client
