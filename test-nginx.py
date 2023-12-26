#!/usr/bin/env python3

import tempfile
import subprocess
import os
import urllib.request
import time

from contextlib import contextmanager

script_path = os.path.realpath(__file__)
script_dir = os.path.dirname(script_path)

@contextmanager
def nginx():
    tmp = tempfile.TemporaryDirectory()
    cfg = os.path.join(script_dir, "doc/nginx.conf")
    os.mkdir(os.path.join(tmp.name, "logs"))
    p = subprocess.Popen(["nginx", "-p", tmp.name, "-c", cfg])
    try:
        yield p
    finally:
        p.terminate()
        p.wait()
        tmp.cleanup()

@contextmanager
def go_pkg_proxy():
    exe = os.path.join(script_dir, "target/go-pkg-proxy")
    modules = os.path.join(script_dir, "doc/go.json")
    p = subprocess.Popen([exe, "-log-level", "DEBUG", "-modules", modules])
    try:
        yield p
    finally:
        p.terminate()
        p.wait()

if __name__ == "__main__":
    with nginx() as n, go_pkg_proxy() as g:
        time.sleep(1)
        rsp = urllib.request.urlopen("http://localhost:7000/go-pkg-proxy?go-get=1")
        assert(rsp.status == 200)
        bs = rsp.read()
        assert(b"go-import" in bs)
