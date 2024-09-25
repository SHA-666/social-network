import Link from 'next/link';
import styles from '../styles/Home.module.css';

export default function Home() {
  return (
    <div className={styles.container}>
      <header className={styles.header}>
        <h1>Bienvenue sur social forum</h1>
        <p>Une plateforme sociale où vous pouvez vous connecter et partager avec vos amis.</p>
      </header>

      <main className={styles.main}>
        <div className={styles.actions}>
        
        </div>
      </main>

      <footer className={styles.footer}>
        <p>&copy; 2024. Tous droits réservés.</p>
      </footer>
    </div>
  );
}