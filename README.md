## DB 

DB is an attempt to prove to myself that Go stdlib is more production ready
and more powerful than python's one.

I often hear people saying that python stdlib is batteries include, I agree but,
what really are theses batteries? As far as I know, these batteries are
just utilities for toy projects, a part from asyncio, functools and iterators
modules, everything else in the stdlib seems not to be production ready.

I am now convinced !

### Do you want to try it ?

It is just a toy project where I played a little bit with Go stdlib `*sql.DB`, so
don't except something exceptional from it.

```bash
$ cd db; make build
$ ./dist/db --datasource=articles.db
```
