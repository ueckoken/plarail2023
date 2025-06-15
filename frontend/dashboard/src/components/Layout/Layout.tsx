import React from 'react';
import styles from './Layout.module.scss';

interface LayoutProps {
  children: React.ReactNode;
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <div className={styles.layout}>
      <header className={styles.header}>
        <div className={styles.headerContent}>
          <div className={styles.logo}>
            <svg width="32" height="32" viewBox="0 0 32 32" fill="none">
              <rect width="32" height="32" rx="6" fill="#1976D2"/>
              <path d="M8 16H24M8 12H24M8 20H24" stroke="white" strokeWidth="2" strokeLinecap="round"/>
              <circle cx="10" cy="16" r="2" fill="#FFC107"/>
              <circle cx="22" cy="16" r="2" fill="#FFC107"/>
            </svg>
            <h1 className={styles.title}>鉄道制御システム</h1>
          </div>
          <nav className={styles.nav}>
            <button className={styles.navButton}>
              <svg width="20" height="20" viewBox="0 0 20 20" fill="currentColor">
                <path d="M10 12a2 2 0 100-4 2 2 0 000 4z"/>
                <path fillRule="evenodd" d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z" clipRule="evenodd"/>
              </svg>
              <span>監視モード</span>
            </button>
            <button className={styles.navButton}>
              <svg width="20" height="20" viewBox="0 0 20 20" fill="currentColor">
                <path fillRule="evenodd" d="M11.49 3.17c-.38-1.56-2.6-1.56-2.98 0a1.532 1.532 0 01-2.286.948c-1.372-.836-2.942.734-2.106 2.106.54.886.061 2.042-.947 2.287-1.561.379-1.561 2.6 0 2.978a1.532 1.532 0 01.947 2.287c-.836 1.372.734 2.942 2.106 2.106a1.532 1.532 0 012.287.947c.379 1.561 2.6 1.561 2.978 0a1.533 1.533 0 012.287-.947c1.372.836 2.942-.734 2.106-2.106a1.533 1.533 0 01.947-2.287c1.561-.379 1.561-2.6 0-2.978a1.532 1.532 0 01-.947-2.287c.836-1.372-.734-2.942-2.106-2.106a1.532 1.532 0 01-2.287-.947zM10 13a3 3 0 100-6 3 3 0 000 6z" clipRule="evenodd"/>
              </svg>
              <span>設定</span>
            </button>
          </nav>
        </div>
      </header>
      
      <main className={styles.main}>
        <aside className={styles.sidebar}>
          <div className={styles.statusCard}>
            <h3>システム状態</h3>
            <div className={styles.statusItem}>
              <span className={styles.statusIndicator} data-status="active"></span>
              <span>正常稼働中</span>
            </div>
          </div>
          
          <div className={styles.controlCard}>
            <h3>制御パネル</h3>
            <button className={styles.controlButton} data-variant="primary">
              <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor">
                <path d="M8 4.754a3.246 3.246 0 100 6.492 3.246 3.246 0 000-6.492zM5.754 8a2.246 2.246 0 114.492 0 2.246 2.246 0 01-4.492 0z"/>
                <path d="M9.796 1.343c-.527-1.79-3.065-1.79-3.592 0l-.094.319a.873.873 0 01-1.255.52l-.292-.16c-1.64-.892-3.433.902-2.54 2.541l.159.292a.873.873 0 01-.52 1.255l-.319.094c-1.79.527-1.79 3.065 0 3.592l.319.094a.873.873 0 01.52 1.255l-.16.292c-.892 1.64.901 3.434 2.541 2.54l.292-.159a.873.873 0 011.255.52l.094.319c.527 1.79 3.065 1.79 3.592 0l.094-.319a.873.873 0 011.255-.52l.292.16c1.64.893 3.434-.902 2.54-2.541l-.159-.292a.873.873 0 01.52-1.255l.319-.094c1.79-.527 1.79-3.065 0-3.592l-.319-.094a.873.873 0 01-.52-1.255l.16-.292c.893-1.64-.902-3.433-2.541-2.54l-.292.159a.873.873 0 01-1.255-.52l-.094-.32z"/>
              </svg>
              全体制御
            </button>
            <button className={styles.controlButton} data-variant="warning">
              <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor">
                <path d="M8 15A7 7 0 118 1a7 7 0 010 14zm0 1A8 8 0 108 0a8 8 0 000 16z"/>
                <path d="M7.002 11a1 1 0 112 0 1 1 0 01-2 0zM7.1 4.995a.905.905 0 111.8 0l-.35 3.507a.552.552 0 01-1.1 0L7.1 4.995z"/>
              </svg>
              緊急停止
            </button>
          </div>
          
          <div className={styles.infoCard}>
            <h3>運行情報</h3>
            <ul className={styles.infoList}>
              <li>
                <span className={styles.label}>稼働列車数:</span>
                <span className={styles.value}>3</span>
              </li>
              <li>
                <span className={styles.label}>信号状態:</span>
                <span className={styles.value}>正常</span>
              </li>
              <li>
                <span className={styles.label}>ポイント切替:</span>
                <span className={styles.value}>2/5</span>
              </li>
            </ul>
          </div>
        </aside>
        
        <div className={styles.content}>
          {children}
        </div>
      </main>
    </div>
  );
};

export default Layout; 
