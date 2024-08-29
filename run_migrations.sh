if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
fi

# Directory containing the migration files
MIGRATIONS_DIR="migrations"

echo "Running migrations in $MIGRATIONS_DIR"

# Run each .sql file in the directory in order
for file in $MIGRATIONS_DIR/*.sql; do
  echo "Running migration: $file"
  psql -d $DB_NAME -U $DB_USER -f "$file"
  
  if [ $? -ne 0 ]; then
    echo "Error running migration: $file"
    exit 1
  fi
done

echo "All migrations executed successfully!"
