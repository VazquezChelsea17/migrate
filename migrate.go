package migrate

// ... existing imports ...

func (m *Migrate) run(migration *Migration, direction int) error {
	// ... existing logic ...
	
	if err != nil {
		// Check if the driver supports transactions and if we are in a transaction
		if m.isTransactionSupported() {
			// If the transaction rolled back, we don't mark as dirty
			return err
		}
		return m.setVersion(migration.Version, true)
	}
	
	// ... existing logic ...
	return nil
}