describe('Cookie()', function() {
  // Test expires manually.
  if (Cookie.enabled()) {
    var now = new Date();
    now.setDate(now.getDate() + 1);
    var i = 0;
    var prefix = function() {
      i++;
      return 'expires_' + i + '_';
    };

    Cookie.setRaw(prefix() + '_dateString', 'A parsable date string', { expires: now.toGMTString() });
    Cookie.setRaw(prefix() + '_dateObject', 'A date object', { expires: now });
    Cookie.setRaw(prefix() + '_numberValue', 'Default unit is day', { expires: 7 });
    Cookie.setRaw(prefix() + '_oneYearLater', 'Suffix char is an uppercase "Y"', { expires: '1Y' });
    Cookie.setRaw(prefix() + '_oneMonthLater', 'Suffix char is an uppercase "M"', { expires: '1M' });
    Cookie.setRaw(prefix() + '_oneDayLater', 'Suffix char is an uppercase "D"', { expires: '1D' });
    Cookie.setRaw(prefix() + '_oneHourLater', 'Suffix char is an lowercase "h"', { expires: '1h' });
    Cookie.setRaw(prefix() + '_tenMinutesLater', 'Suffix char is an lowercase "m"', { expires: '10m' });
    Cookie.setRaw(prefix() + '_thirtySecondsLater', 'Suffix char is an lowercase "s"', { expires: '30s' });
  };

  it('should return null if no parameter', function() {
    expect(Cookie()).toBe(null);
  });

  it('should return the cookie value if only one parameter', function() {
    document.cookie = 'only_one=value';
    expect(Cookie('only_one')).toBe('value');
  });

  it('should remove the cookie value if the second parameter is null', function() {
    document.cookie = 'second_null=value;path=/';
    expect(Cookie('second_null')).toBe('value');

    Cookie('second_null', null);
    expect(Cookie('second_null')).toBe(null);
  });

  it('should return the set cookie value if the second parater is not null', function() {
    Cookie('one', '1');
    expect(Cookie('one')).toBe('1');
  });
});

describe('Cookie.enabled()', function() {
  it('should return true if the cookie is enabled', function() {
    expect(Cookie.enabled()).toBe(true);
  });

  it('should return false if the cookie is disabled', function() {
    expect(Cookie.enabled()).toBe(false);
  });
});

describe('Cookie.get()', function() {
  it('should return null if the cookie key isn\'t exist', function() {
    expect(Cookie.get('nokey')).toBe(null);
  });

  it('should return null if no parameter passed', function() {
    expect(Cookie.get()).toBe(null);
  });

  it('should return "Bar" if "foo" passed', function() {
    document.cookie = 'foo=Bar';
    document.cookie = 'foofoo=BarBar';
    expect(Cookie.get('foo')).toBe('Bar');
  });

  it('should return the decoded string if cookie value is encoded', function() {
    var homepage = 'https://github.com/Alex1990/tiny-cookie';
    document.cookie = 'homepage=' + encodeURIComponent(homepage);
    expect(Cookie.get('homepage')).toBe(homepage);
  });

  it('should return the value without decoding if second parameter is true', function() {
    var key = 'withoutDecoding';
    var value = encodeURIComponent('https://github.com/Alex1990/tiny-cookie');
    document.cookie = key + '=' + value;
    expect(Cookie.get(key, true)).toBe(value);
  });

  it('should return right value if cookie key contains whitespace', function() {
    var key = 'he  llo';
    var value = 'world';
    document.cookie = key + '=' + value;
    expect(Cookie.get(key)).toBe(value);
  });

  it('should return right value if cookie value contains whitespace', function() {
    var key = 'whitespacevalue';
    var value = 'va  lue';
    document.cookie = key + '=' + value;
    expect(Cookie.get(key)).toBe(value);
  });

  it('should return null if empty string passed', function() {
    document.cookie = 'onlyvalue';
    expect(Cookie.get('')).toBe(null);

    document.cookie = '=tiny=cookie';
    expect(Cookie.get('')).toBe(null);
  });

  // At least, in IE 9, it is impossible that cookie value is an empty string.
  it('should return an empty string if only key is set', function() {
    document.cookie = 'onlykey=';
    expect(Cookie.get('onlykey')).toBe('');
  });
});

describe('Cookie.getRaw()', function() {
  it('should return the value without decoding', function() {
    var key = 'withoutDecoding';
    var value = encodeURIComponent('https://github.com/Alex1990/tiny-cookie');
    document.cookie = key + '=' + value;
    expect(Cookie.getRaw(key)).toBe(value);
  });
});

describe('Cookie.set()', function() {
  it('should return the set cookie value', function() {
    Cookie.set('someKey', 'someValue');
    expect(Cookie.get('someKey')).toBe('someValue');
  });

  it('should return the decoded value', function() {
    var github = 'https://github.com/Alex1990';
    Cookie.set('github', github);
    expect(Cookie.get('github')).toBe(github);
  });

  it('should return the value without encoding if the third parameter is true', function() {
    var key = 'withoutEncoding';
    var value = 'https://github.com/Alex1990/tiny-cookie';
    Cookie.setRaw(key, value);
    expect(Cookie.getRaw(key)).toBe(value);
  });

  it('should return null when cookie path is restricted', function() {
    Cookie.set('path_cookie', 'some_value', { path: '/the-other-path/' });
    expect(Cookie.get('path_cookie')).toBe(null);
  });

  it('should set secure when option is true', function() {
    Cookie.set('someKey', 'someValue', { secure: true });
    expect(Cookie.get('someKey')).toBe(null);
  });

  it('should not set secure when option is false', function() {
    Cookie.set('someKey', 'someValue', { secure: false });
    expect(Cookie.get('someKey')).toBe('someValue');
  });
});

describe('Cookie.setRaw()', function() {
  it('should return the value without encoding', function() {
    var key = 'withoutEncoding';
    var value = 'https://github.com/Alex1990/tiny-cookie';
    Cookie.setRaw(key, value);
    expect(Cookie.getRaw(key)).toBe(value);
  });
});

describe('Cookie.remove()', function() {
  it('should return null when remove "removeKey" cookie', function() {
    Cookie.set('removeKey', 'removeValue');
    Cookie.remove('removeKey');
    expect(Cookie.get('removeKey')).toBe(null);
  });

  // it('should return null whe remove the empty string key', function() {
  //   Cookie.remove('');
  //   expect(Cookie.get('')).toBe(null);
  // });
});
