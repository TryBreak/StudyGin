window.api = (function () {
  const main = (data) => {
    return http.ajax({
      url: '/',
      data,
      method: 'get',
    });
  };

  const admin = (data) => {
    return http.ajax({
      url: '/admin',
      data,
      method: 'post',
    });
  };

  const ping = (data) => {
    return http.ajax({
      url: '/ping',
      data,
      method: 'get',
    });
  };

  const user_name = (data) => {
    return http.ajax({
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
