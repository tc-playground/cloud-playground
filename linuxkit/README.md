# LinuxKit

## Introduction

* `LinuxKit`, is a toolkit for building custom minimal, immutable Linux distributions based on `containers`.

* `LinuxKit` uses the `linuxkit` tool for `building`, `pushing` and `running` `VM images`.

* `LinuxKit` uses a `linuxkit.yml` file to `define` the `image` to `build`.

* `LinuxKit` supports a set of `platforms` for `running` images on. These can be `local`, `cloud`, or `baremetal` based.

* `LinuxKit Images` contain a base `kernel`, `init system`, and, set of `containerised services`.

    > NB: `containerd` is used to run the containers.


---

## linuxkit.yml - Image specification.

* [`linuxkit.yml`](https://github.com/linuxkit/linuxkit/blob/master/linuxkit.yml) - Defines the `image` to build.


* `linuxkit.yml` specifies a set of `resources` that are built into the generated `image` and started at boot time. 

    * A `kernel`
    
    * A base `init system`. 
    
    * A set of `containerised services`.

* `linuxkit.yml` documentation can be found [here](https://github.com/linuxkit/linuxkit/blob/master/docs/yaml.md).

* `linuxkit.yml` examples can be found [here](https://github.com/linuxkit/linuxkit/tree/master/examples).


* `linuxkit.yml` resources are specified as Docker `images` which are pulled from a `registry`.

    * Optionally, verified with `Docker Content Trust`.

    * Private registries require a `docker login`.

* `linuxkit.yml` containers will be run in the host `net`, `ipc` and `uts` namespaces by default.

* `linuxkit.yml` images are run sequentially and each must exit before the next one is run.

* `linuxkit.yml` `mount point`s must already exist, as must a file or directory being bind mounted into a container.


> NB: All images must confirm to the[OCI Specification](https://github.com/opencontainers/runtime-spec/blob/master/spec.md).

> NB:  `linuxkit.yml` documentation can be found [here](https://github.com/linuxkit/linuxkit/blob/master/docs/yaml.md).

---

### `linuxkit.yml` - top-level elements

1. __kernel__ - Defines the kernel configuration. The files will be put into the `boot/` directory, where they are used to build `bootable images`. Required if booting a VM.

2. __init__ - Defines a list of OCI images that are used for the `init system` and are `unpacked directly into the root filesystem`.

    > NB: This should bring up `containerd`, start the `system` and `daemon` containers, and set up basic filesystem `mounts`. 

3. __onboot__ -  Defines a list of OCI images to run on `startup`.

4. __onshutdown__ - Defines a list of OCI images to run on a clean `shutdown`.

5. __services__ - Defines a list of OCI images for long running `services` which are run with `containerd`.

6. __files__ - Adds files inline in the config, or from an external file.

7. __trust__ - Specifies which build components are to be cryptographically verified with Docker Content Trust prior to pulling.

---

## `Platforms` - Image runtime `Hypervisors` 

* `LinuxKit` images can be run on a supported `platform hypervisor`.

    * `Local Hypervisors`

        * `qemu` - macOS, Linux, Windows - (x86_64, arm64, s390x)
        * `HyperKit` - macOS - (x86_64)
        * `Hyper-V` - Windows - (x86_64)
        * `VMware` - macOS, Windows - (x86_64)
    
    * `Cloud based platforms`

        * `Amazon Web Services` - (x86_64)
        * `Google Cloud` - (x86_64)
        * `Microsoft Azure` - (x86_64)
        * `OpenStack` - (x86_64)
        * `Scaleway` - (x86_64)
    
    * `Baremetal`

        * `packet.net`- ([x86_64, arm64)
        * `Raspberry Pi Model 3b` [arm64]


---

## Commands

### `Build` Commands

* The `linuxkit build` command assembles `a set of containerised components` into in `image`.

* `linuxkit build -help` - Get help on `building` an image.

* `linuxkit build linuxkit.yml` - Build a new image.

    * `linuxkit build -format raw-bios linuxkit.yml` - to output a `raw BIOS` bootable disk image, 
    
    * `linuxkit build -format iso-efi linuxkit.yml` - to output an `EFI` bootable ISO image.

    * `linuxkit build -format vhd linuxkit.yml` - to output an `VHD` bootable ISO image


### `Run Commands`

* `linuxkit run --help` - Get help on `running` and image.

* `linuxkit run <name>` or `linuxkit run <name>.<format>` to execute the image you created with `linuxkit build <name>.yml`. 

    * This will use a `suitable backend` for your platform or you can choose one, for example `qemu`. 



---

## References

* [LinuxKit Home](https://github.com/linuxkit/linuxkit)

* [LinuxKit Advantages](https://collabnix.com/top-10-reasons-why-linuxkit-is-better-than-the-traditional-os-distribution/)