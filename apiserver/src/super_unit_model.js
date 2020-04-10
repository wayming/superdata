const Pool = require('pg').Pool
const pool = new Pool({
  user: 'postgres',
  host: 'db',
  database: 'dev',
  password: 'postgres',
  port: 5432,
});


const getUnits = () => {
    return new Promise(function(resolve, reject) {
      pool.query('SELECT * FROM SUNSUPER ORDER BY vdate ASC', (error, results) => {
        if (error) {
          reject(error)
        }
        resolve([ {key: 'SUN', values: results.rows} ]);
      })
    }) 
  }

  module.exports = {
    getUnits,
  }