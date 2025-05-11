import streamlit as st
import requests
import json

st.title("Stock Price Checker")

stock_name = st.text_input("Enter stock name")
start_date = st.date_input("Enter start date")
end_date = st.date_input("Enter end date")

if st.button("Get Results"):
    data = {
        "stockName": stock_name,
        "startDate": str(start_date),
        "endDate": str(end_date),
    }

    # response = requests.get("http://localhost:8080/getdata", data)
    # response = requests.get("http://172.17.0.2:8080/getdata", data)
    response = requests.get("http://go-server-service.default.svc.cluster.local:8080/getdata", data)

    if response.status_code == 200:
        print(response.text)
        data = json.loads(response.text)
        st.write("Results: ")
        st.write("Max price and date: ", data[0]['Date'][:10], data[0]['Price'])
        st.write("Min price and date: ", data[1]['Date'][:10], data[1]['Price'])

    else:
        st.write("Error in fetching from backend")