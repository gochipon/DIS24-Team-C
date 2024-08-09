import psycopg2
import json

class SQLquery:
    
    def __init__(self, db_url, records):
        self.db_url = db_url
        self.records = records
        
    def generate_query(self, records):
        query_list = []
        for record in records:
            if record["name"] == "issues":
                query = f"SELECT body, title, url FROM issues WHERE number = CAST({record['id']} as BIGINT);"
            elif record["name"] == "pull_requests":
                query = f"SELECT body, title, url FROM pull_requests WHERE number = CAST({record['id']} as BIGINT);"
            elif record["name"] == "releases":
                query = f"SELECT body, title, url FROM releases WHERE name = \'{record['id']}\' AND repository = \'{record['repository']}\';"    
            query_list.append(query)
        return query_list
        
    def execute(self):
        query_list = self.generate_query(self.records)
        results = []
        connection = psycopg2.connect(self.db_url)
        cursor = connection.cursor()
        for query in query_list:
            try:
                cursor.execute(query)
                result = cursor.fetchall()
                connection.commit()
                results.append({"body": result[0][0], "title": result[0][1], "url": result[0][2]})
            except (Exception, psycopg2.Error) as error:
                print(f"Error while connecting to PostgreSQL: {error}")
        if connection:
            cursor.close()
            connection.close()
        return results
        
def main(records: list) -> list:
    config = json.load(open('config.json'))
    DATABASE_URL = config['DATABASE_URL']
    sql_query = SQLquery(DATABASE_URL, records)
    results = sql_query.execute()
    return results