
## tzpipe utility for minibus message handling from the shell

_tzpipe_ is a command line utility that lets you send and receive messages
to and from a minibus service.

## USAGE 

_tzpipe_ _[send|listen]_ _channel_ arg1 arg2..


### Listening:

tzpipe always needs a verb (send | listen)  and a minibus channel name to operate on, 
with no other parameters this listens on the selected channel and outputs to `STDOUT`. 

Here we connect to a ficticious channel that is being sent random animals:
```sh
> tzpipe listen animals
cats
dogs
rabbits
...
```

We can give tzpipe a regular expression (using Golang's regex syntax)
```sh
> tzpipe listen animals --filter='^dogs'
dogs
dogs
dogs 
...
```

While this is interesting, sometimes you may want to wait for a particular message
and then act on it in a script, you can do this by using the --limit argument which 
will exit tzpipe after a given number of messages:
```sh
l=$(tzpipe listen animals --filter='llama' --limit=1)
echo "Found a $l!"
```

### Sending

tzpipe allows you to send messages to channels as well, there are a number of
ways to do this:

```sh
> tzpipe send foo 'This is a message'
```

tzpipe accepts `STDIN` and will send each line to the channel, you can use
--limit here as well to stop capturing after a number of lines:
```sh
cat /dev/poetry | tzpipe send foo --limit=100
```

You can use tzpipe to reduce incomming messages and deliver them to a different
channel, here we filter cats from the animals channel, log matches to `STDERR` and 
send them on to the 'catsonly' channel:
```sh
tzpipe listen animals --filter='^cats' --log | tzpipe send catsonly
```

<br/><br/>
<table>
<tr><td>
<img src='https://raw.githubusercontent.com/ttyzero/logo/master/assets/ttyzero_animated.png' alt='ttyZero Logo' title='ttyZero Logo'/>
</td>
<td style='padding-left: 10em'>
<h2>tzpipe is part of the <a href='http://github.com/ttyzero'>ttyZero Project</a></h2>
<b>tzpipe</b> is <i>(c) 2018 ttyZero authors</i> <br/>
 and is available under the <b>MIT license</b>. 
</td></tr>
</table>
