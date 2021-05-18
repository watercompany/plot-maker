## plot-maker

Run ProofOfSpace with arguments from a JSON file.

```
Usage of plot-maker:
  -2 string
        Override value for tmp2_dir
  -b int
        Override value for buffer
  -bin string
        ProofOfSpace binary path (default "ProofOfSpace")
  -d string
        Override value for final_dir
  -debug
        Debug parameters
  -f string
        Override value for filename
  -i string
        Override value for plot_id
  -json string
        Input JSON
  -k int
        Override value for size
  -m string
        Override value for plot_memo
  -r int
        Override value for num_threads
  -s int
        Override value for stripe_size
  -t string
        Override value for tmp_dir
  -version
        Prints versin and exit
```

## Building

```
make build
```

## Running

- All parameters coming from the JSON config file can be overriden using the same short form that `ProofOfSpace` CLI uses.
- `ProofOfSpace` CLI must be present in `$PATH` or specified using the `-bin` flag.

### Example config
```json
{
    "tmp_dir": ".",
    "tmp2_dir": ".",
    "final_dir": ".",
    "filename": "plot-k25-2021-05-17-20-13-b22313297558bd1ad8414e5f6d24bdc3163430c1a45b0fd0d97c802c10241513.plot",
    "size": 25,
    "plot_memo": "aac13cb5fb04e85d1a26ffe5c155b3089f527af3574621bb4d59d364fd628a0cc80edc042c9a8139e2b62c15398259ac89b692b9769337f370a483d72dc9529ac96d2162ae1ceb51d152063681a044fdd01168b91e4bf53dc6e7ba0f8bc34bdc1f5be03430256c843f6078c554bddf22b600a9dc96e264f3da0454b41ff1b70e",
    "plot_id": "b22313297558bd1ad8414e5f6d24bdc3163430c1a45b0fd0d97c802c10241513",
    "buffer": 3389,
    "stripe_size": 65536,
    "num_threads": 24,
    "nobitfield": false
}
```

```
./plot-maker -json example-args.json
```

### Running with Docker

```bash
make docker

docker run --rm -it \
    -v $(pwd):/root/final_dir \
    -v $(pwd)/example-args.json:/root/args.json \
    plot-maker:latest
```
