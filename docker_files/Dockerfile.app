FROM python:3.9

WORKDIR /app

COPY frontend .

RUN pip install streamlit

EXPOSE 8501

CMD ["streamlit", "run", "app.py"]