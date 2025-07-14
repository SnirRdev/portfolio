FROM python:3.12-slim

RUN addgroup -S appuser && adduser -S appuser -G appuser

WORKDIR /app

COPY --chown=appuser:appuser ./app /app
RUN pip install --no-cache-dir -r requirements.txt

USER appuser

EXPOSE 5000
CMD ["gunicorn", "--bind", "0.0.0.0:5000", "app:app"]


