# tiny-cookie

A tiny cookie manipulation plugin.

The tiny-cookie will expose a method `Cookie` on the global scope. Also, it can be as a CommonJS/AMD module.

## Packages

**NPM:**

```bash
npm install tiny-cookie
```

**Bower:**

```bash
bower install tiny-cookie
```

## APIs

### Cookie.enabled()

Check if the cookie is enabled.

### Cookie.get(key)

**Alias: Cookie(key)**

Get the cookie value with decoding, using `decodeURIComponent`.

### Cookie.getRaw(key)

**Also: Cookie.get(key, true)**

Get the cookie value without decoding.

### Cookie.set(key, value, options)

**Alias: Cookie(key, value, options)**

Set a cookie with encoding the value, using `encodeURIComponent`. The `options` parameter is an object. And its property can be a valid cookie option, such as `path`(default: root path `/`), `domain`, `expires`/`max-age` or `secure` (Note: the `secure` flag will be set if it is an truthy value, such as `true`, or it will be not set). For example, you can set the expiration:

```js
var now = new Date;
now.setMonth(now.getMonth() + 1);

Cookie.set('foo', 'Foo', { expires: now.toGMTString() });
```

The `expires` property value can accept a `Date` object, a parsable date string (parsed by `Date.parse()`), an integer (unit: day) or a numeric string with a suffix character which specifies the time unit.

| Unit suffix | Representation |
| ----------- | -------------- |
| Y           | One year       |
| M           | One month      |
| D           | One day        |
| h           | One hour       |
| m           | One minute     |
| s           | One second     |

**Examples:**

```js
var date = new Date;
date.setDate(date.getDate() + 21);

Cookie.set('dateObject', 'A date object', { expires: date });
Cookie.set('dateString', 'A parsable date string', { expires: date.toGMTString() });
Cookie.set('integer', 'Seven days later', { expires: 7 });
Cookie.set('stringSuffixY', 'One year later', { expires: '1Y' });
Cookie.set('stringSuffixM', 'One month later', { expires: '1M' });
Cookie.set('stringSuffixD', 'One day later', { expires: '1D' });
Cookie.set('stringSuffixh', 'One hour later', { expires: '1h' });
Cookie.set('stringSuffixm', 'Ten minutes later', { expires: '10m' });
Cookie.set('stringSuffixs', 'Thirty seconds later', { expires: '30s' });
```

### Cookie.setRaw(key, value, options)

**Also: Cookie.set(key, value, true, options)**

Set a cookie without encoding.

### Cookie.remove(key)

**Alias: Cookie(key, null)**

Remove a cookie.

## License

MIT.
