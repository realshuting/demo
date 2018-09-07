const count = db.getMongo().getDBNames().reduce((total, database) => {
  if(database != 'config' && database != 'local') {
    const dbMongo = db.getSiblingDB(database)
    total += dbMongo.getCollectionNames().reduce((count, col) => {
      count += dbMongo.getCollection(col).find().length()
      return count
    }, 0)
  }
  return total
}, 0)
print('Total number of records is', count);
