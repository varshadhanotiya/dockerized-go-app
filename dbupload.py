import pandas as pd
from sqlalchemy import create_engine

# Function to import data from a CSV file into a table
def import_csv_to_table(table_name, csv_file_path):
    df = pd.read_csv(csv_file_path)
    engine = create_engine('postgresql://postgres:postgres@127.0.0.1:5432/nifty')
    df.to_sql(table_name, engine)

def main():
    path1 = "C:\\Users\\varsha\\Documents\\GitHub\\golang-devproject\\data\\Final-50-stocks.csv"
    import_csv_to_table("summary", path1)

if __name__ == '__main__':
    main()
