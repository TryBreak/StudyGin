window.api = (function () {
  const main = (data) => {
    http.ajax({
      url: '/',
      data: {
        abc: 123,
      },
      method: 'get',
    });
  };

  const admin = (data) => {
    http.ajax({
      url: '/admin',
      data: {
        abc: 123,
      },
      method: 'post',
    });
  };

  const ping = (data) => {
    http.ajax({
      url: '/ping',
      data: {
        abc: 123,
      },
      method: 'get',
    });
  };

  const user_name = (data) => {
    http.ajax({
      url: `/user/${data.name}`,
      data,
      method: 'get',
    });
  };

  return {
    main,
    admin,
    ping,
    user_name,
  };
})();
